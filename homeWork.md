## 1)完善solidity代码与生成的go代码

> ##  ***solidity代码:***

```solidity
pragma solidity ^0.6.0;

contract UserManager {
    address payable public owner;
    
    mapping(uint8 => string) accounts;    //id -> passwd
    mapping(uint8 => address) ips;       //IP address
    
    event Login(uint8 id, uint time);
    event Register(uint8 id ,string passwd);
    event SetPassword(string passwd,uint time);
    
    constructor () public {
        owner = msg.sender;
    }

    function login(uint8 id, string memory passwd) public returns (bool) {
        require(ips[id] == msg.sender);
        //这里用了哈希值的比对,因为两个数组是不能直接用==号来进行比对的
        //所以先用bytes()转换了这两个数组的数据类型，在用keccak256()来转换成对应的哈希值
        if (keccak256(bytes(accounts[id])) == keccak256(bytes(passwd))) {
            emit Login(id, now);
            return true; 
        }
        return false;
    }
    
    function getIP(uint8 id) public view returns (address) {
        require(ips[id] != address(0));
        return ips[id];
    }
    
    function register(uint8 id,string memory passwd) public returns (bool) {
        if(keccak256(bytes(accounts[id]))!=''){
        //这里是判断id有没有被使用
            accounts[id]=passwd;
        //这里是让id和设置的密码进行映射
            emit  Register(id,passwd);
        //这里是触发了事件并传回了id和密码
        return true;
        }else{
        //如果id被使用就会返回false
            return false;
        }
        
    }
    
    function setPassword(uint8 id,string memory Usedpasswd,string memory passwd) public returns (bool) {
    	//这是判断旧的密码是否正确
         if (keccak256(bytes(accounts[id])) == keccak256(bytes(Usedpasswd))) {
        //这是修改密码，将id对应的映射改为传入的新密码
            accounts[id]=passwd;
            emit SetPassword(passwd,now);
        //触发了事件返回了新密码和时间
            return true;
        }else{
        //旧密码不对就返回false
            return false;
        }
        
    }
  }
```

> ## 心得:
>
> * 明白了storage和memory的存储关系，熟悉了怎么将storage转换成memory、memory转换成storage、memory转换成memory、以及storage转换成storage。
>
> * 了解了solitdity里面事件的使用方式和命名方式。.
>
> * 知道了go mod对soildity的重要性，必须要有mod库才能用gcc后的go文件不报错。

## 2)阅读文档：

> ## solidity代码:

```solidity
pragma solidity ^0.6.0;

/**
 * @title CheckpointOracle
 * @author Gary Rong<garyrong@ethereum.org>, Martin Swende <martin.swende@ethereum.org>
 * @dev Implementation of the blockchain checkpoint registrar.
 */
contract CheckpointOracle {
    /*
        Events
    */

    // NewCheckpointVote is emitted when a new checkpoint proposal receives a vote.
    event NewCheckpointVote(uint64 indexed index, bytes32 checkpointHash, uint8 v, bytes32 r, bytes32 s);

    /*
        Public Functions
    */
    //初始化赋值admins和adminList
    constructor(address[] memory _adminlist, uint _sectionSize, uint _processConfirms, uint _threshold) public {
        for (uint i = 0; i < _adminlist.length; i++) {
            admins[_adminlist[i]] = true;
            adminList.push(_adminlist[i]);
        }
        sectionSize = _sectionSize;
        processConfirms = _processConfirms;
        threshold = _threshold;
    }

    /**
     * @dev Get latest stable checkpoint information.
     * @return section index
     * @return checkpoint hash
     * @return block height associated with checkpoint
     */
     //获取到最后一个节点
    function GetLatestCheckpoint()
    view
    public
    returns(uint64, bytes32, uint) {
        return (sectionIndex, hash, height);
    }

    // SetCheckpoint sets  a new checkpoint. It accepts a list of signatures
    // @_recentNumber: a recent blocknumber, for replay protection
    // @_recentHash : the hash of `_recentNumber`
    // @_hash : the hash to set at _sectionIndex
    // @_sectionIndex : the section index to set
    // @v : the list of v-values
    // @r : the list or r-values
    // @s : the list of s-values
    //设置节点
    function SetCheckpoint(
        uint _recentNumber,
        bytes32 _recentHash,
        bytes32 _hash,
        uint64 _sectionIndex,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s)
        public
        returns (bool)
    {
        // Ensure the sender is authorized.
        require(admins[msg.sender]);

        // These checks replay protection, so it cannot be replayed on forks,
        // accidentally or intentionally
        require(blockhash(_recentNumber) == _recentHash);

        // Ensure the batch of signatures are valid.
        require(v.length == r.length);
        require(v.length == s.length);

        // Filter out "future" checkpoint.
        //判断节点是否比区块链的块多
        if (block.number < (_sectionIndex+1)*sectionSize+processConfirms) {
            return false;
        }
        // Filter out "old" announcement
        
        if (_sectionIndex < sectionIndex) {
            return false;
        }
        // Filter out "stale" announcement
        if (_sectionIndex == sectionIndex && (_sectionIndex != 0 || height != 0)) {
            return false;
        }
        // Filter out "invalid" announcement
        if (_hash == ""){
            return false;
        }

        // EIP 191 style signatures
        //
        // Arguments when calculating hash to validate
        // 1: byte(0x19) - the initial 0x19 byte
        // 2: byte(0) - the version byte (data with intended validator)
        // 3: this - the validator address
        // --  Application specific data
        // 4 : checkpoint section_index(uint64)
        // 5 : checkpoint hash (bytes32)
        //     hash = keccak256(checkpoint_index, section_head, cht_root, bloom_root)
        bytes32 signedHash = keccak256(abi.encodePacked(byte(0x19), byte(0), this, _sectionIndex, _hash));

        address lastVoter = address(0);

        // In order for us not to have to maintain a mapping of who has already
        // voted, and we don't want to count a vote twice, the signatures must
        // be submitted in strict ordering.
        for (uint idx = 0; idx < v.length; idx++){
            address signer = ecrecover(signedHash, v[idx], r[idx], s[idx]);
            require(admins[signer]);
            require(uint256(signer) > uint256(lastVoter));
            lastVoter = signer;
            emit NewCheckpointVote(_sectionIndex, _hash, v[idx], r[idx], s[idx]);

            // Sufficient signatures present, update latest checkpoint.
            if (idx+1 >= threshold){
                hash = _hash;
                height = block.number;
                sectionIndex = _sectionIndex;
                return true;
            }
        }
        // We shouldn't wind up here, reverting un-emits the events
        revert();
    }

    /**
     * @dev Get all admin addresses
     * @return address list
     */
    function GetAllAdmin()
    public
    view
    returns(address[] memory)
    {
        address[] memory ret = new address[](adminList.length);
        for (uint i = 0; i < adminList.length; i++) {
            ret[i] = adminList[i];
        }
        return ret;
    }

    /*
        Fields
    */
    // A map of admin users who have the permission to update CHT and bloom Trie root
    mapping(address => bool) admins;

    // A list of admin users so that we can obtain all admin users.
    address[] adminList;

    // Latest stored section id
    uint64 sectionIndex;

    // The block height associated with latest registered checkpoint.
    uint height;

    // The hash of latest registered checkpoint.
    bytes32 hash;

    // The frequency for creating a checkpoint
    //
    // The default value should be the same as the checkpoint size(32768) in the ethereum.
    uint sectionSize;

    // The number of confirmations needed before a checkpoint can be registered.
    // We have to make sure the checkpoint registered will not be invalid due to
    // chain reorg.
    //
    // The default value should be the same as the checkpoint process confirmations(256)
    // in the ethereum.
    uint processConfirms;

    // The required signatures to finalize a stable checkpoint.
    uint threshold;
}
```

>该智能合约设置了节点，并能获取到所有的节点，和获取到最后的一个节点。