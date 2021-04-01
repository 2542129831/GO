# 完善数字身份智能合约

```
pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

contract PersonDID {
    struct Person {
        uint8 id;
        uint8 age;
        string name;
        string story;
    }
    //判断是否为管理员
    modifier Onlyadmin(){
        require(
           isPersonExsist[msg.sender] == true
            ,"You dont'rhave right to dp that!");
        _;
    }
    mapping(uint8 => bool) public Right;
    event AddPerson(uint8 id, uint8 age, string name, uint timestamp);
    event AddStory(uint8 id,string story,uint timestamp);
    //添加个人经历事件
    event giveAccess(uint8 id,uint8 ip,uint timestamp);
   	//添加给予权限的事件
   	Person admin;
    Person[] persons;
    mapping(address => Person) public PersonInfo;
    mapping(address=> bool) public isPersonExsist;
    
    constructor (address ip, uint8 id, string memory name, uint8 age) public {
        admin = Person(id, age, name,"");
        Person memory p = Person(id, age, name,"");
        persons.push(p);
        PersonInfo[msg.sender] = p;
        isPersonExsist[msg.sender] = true;
        Right[id]=false;
    }
    function getNumberOfPersons() view public returns (uint256) {
        return persons.length;
    }
    //只允许管理员来操作添加人员和人员的各项信息
    function addPerson(uint8 id, uint8 age, string memory name,string memory story) Onlyadmin public returns (bool) {
        require(!((id == 0) || age == 0), "persons info can not be empty!!");
        Person memory person = Person(id, age, name,story);
        persons.push(person);
        PersonName[id]=persons[id];
        PersonInfo[msg.sender] = Person(id, age, name,story);
        Right[id]=false;
        emit AddPerson(id, age, name, now);
        emit AddStory(id,story,now);
        return true;
    }
    
    function setPersonAgeSto(address ip, uint8 age) public {
        Person storage p = PersonInfo[ip];
        p.age = age;
    }
    
    function setPersonAgeMem(address ip, uint8 age) public {
        Person memory p = PersonInfo[ip];
        p.age = age;
    }
    
    function getPersonAge(address ip) public view returns (uint8) {
        return PersonInfo[ip].age;
    }
    function setAdminer(address ip) Onlyadmin public returns(bool){
        isPersonExsist[ip]=true;
        return true;
    }
    // string nove;
    // function setStory(uint8 id ,string memory storys)public returns (string memory){
    //     nove=storys;
    //     if(Right[id]){
    //         persons[id].story=nove;

    //     }else
    //     return nove;
    // }
    // function setRight(uint8 id , bool right) Onlyadmin public returns(bool){
    //     Right[id]=right;
    //     return true;
    // }
    //新建一个自己id和别人id关联的映射
    mapping(uint8 =>mapping(uint8 => bool)) PersonRight;
    //新建一个自己id和自己各项属性的映射
    mapping(uint8 =>Person) public PersonName;
    //允许访问次数
    uint8 times;
    //给予其他id可以查看自己各项值的权限，和可以访问的次数
    function giveRight(uint8 Yourid,uint8 ip,uint8 time)public returns (bool){
        PersonRight[Yourid][ip]=true;
        times=time;
        emit giveAccess(Yourid,ip,now);
        return true;
    }
    //查看其他人的信息
    function myself(uint8 seeId,uint8 Yourid)  public returns (Person memory){
    //判断是否给你了权限
        require(PersonRight[seeId][Yourid]);
        if(times==0){
            revert("sorry you times is running out!");
        }else{
            times--;
            return PersonName[seeId];
        }
    }
        
}
```

