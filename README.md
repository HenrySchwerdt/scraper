# Scraper



# Usage

```sh
scraper c --in="input.txt" --out="output.csv"
```

## Where Input.txt looks like this:
```txt
https://www.bild.de/
https://www.nytimes.com/
https://www.bunte.de/
```

## The output will look like this

 
```csv
https://www.nytimes.com/,amp;66,the;50,a;41,of;39,min;36,read;34,the;33,new;32,in;28,and;23,img;22,for;19,york;19,to;18,times;18,delhi;11,opinion;11,with;11,on;10,his;9,a;9,at;8,atul;8,party;8,by;8,congress;8,loke;8,is;7,that;7,audio;6,trump;6,election;6,national;6,news;6,are;6,an;5,advertisement;5,india;5,election;5,in;5,president;5,headquarters;5,skip;5,advertisement;5,indian;5,be;4,et;4,supporters;4,more;4,arts;4,man;4,via;4,results;4,party;4,live;4,june;4,can;4,what;4,prime;4,first;4,narendra;4,center;4,as;4,stop;4,is;4,minister;4,magazine;4,u.s;4,biden;4,has;4,counting;4,its;4,video;4,visual;4,donald;4,israel;3,that;3,modi;3,well;3,verdict;3,will;3,world;3,why;3,lifestyle;3,site;3,it;3,culture;3,words;3,book;3,an;3,bharatiya;3,outside;3,janata;3,seats;3,ukraine;3,he;3,this;3,has;3,i;3,t;3,letters;3,his;3,allies;3,live;3,you;3,more;3,hunter;3,home;3,illustration;3,mexico;3,athletic;3,games;3,since;3,sports;3,nyt;3,was;3,son;2,general;2,alive;2,style;2,case;2,african;2,restaurant;2,have;2,must;2,from;2,network;2,him;2,we;2,jobs;2,manuel;2,reviews;2,nigel;2,two;2,france-presse;2,listen;2,quot;2,as;2,connections;2,quiz;2,trending;2,lópez;2,krugman;2,them;2,food;2,events;2,brand;2,weather;2,new;2,canada;2,first;2,been;2,block;2,michelle;2,trial;2,education;2,farage;2,reuters;2,our;2,privacy;2,terms;2,kendi;2,page;2,obituaries;2,tv;2,its;2,coverage;2,at;2,results;2,wirecutter;2,list;2,reader;2,books;2,through;2,said;2,op-docs;2,essays;2,top;2,family;2,dowd;2,seen;2,early;2,skip;2,politics;2,building;2,our;2,dance;2,andrés;2,show;2,headway;2,fashion;2,health;2,after;2,war;2,columnists;2,ineducation;2,john;2,music;2,how;2,four;2,woman;2,so;2,morning;2,stories;2,pop;2,travel;2,timesmachine;2,international;2,race;2,word;2,when;2,arun;2,be;2,melissa;2,found;2,today;2,learning;2,sellers;2,great;2,changed;2,estate;2,science;2,destruction;2,cooking;2,about;2,majority;2,editorials;2,getty;2,americans;2,my;2,group;2,guest;2,movies;2,puzzle;2,business;2,school;2,real;2,who;2,maureen;2,paul;2,graphics;2,yellow;2,vote;2,television;2,holding;2,killed;2,him;2,love;2,investigations;2,images;2,you;2,also;2,ibram;2,ani;2,goldberg;2,sunday;2,read;2,best;2,bolger;2,photos;2,iranian;2,subscriptions;2,center;2,over;2,corrections;2,theater;2,long;2,obrador;2,tech;2,writes;1,who;1,surrounded;1,power;1,white;1,frank;1,started;1,missed;1,jury;1,fought;1,israeli;1,prosecution;1,via;1,case;1,newshound;1,renkl;1,shrinking;1,felon;1,faces;1,account;1,accessibility;1,choices;1,taken;1,shutterstock;1,india;1,accused;1,when;1,widely;1,where;1,deepfakes;1,use;1,columnist;1,entire;1,virtually;1,ruling;1,parties;1,reckoning;1,newsletter;1,leak;1,which;1,selection;1,personality;1,ammunition;1,than;1,content;1,supreme;1,sins;1,wife;1,navigating;1,dire;1,bankruptcy;1,term;1,their;1,before;1,guards;1,envelopes;1,licensing;1,bordered;1,react;1,rap;1,washington;1,email;1,historic;1,industry;1,moment;1,fears;1,wayne;1,job;1,got;1,syria;1,dark;1,store;1,your;1,using;1,jäger;1,nights;1,far;1,own;1,during;1,were;1,years;1,factory;1,smoke;1,number;1,felt;1,headlines;1,chinese;1,look;1,bee;1,posts;1,company;1,masks;1,wordplay;1,time;1,learn;1,continue;1,painting;1,adrian;1,s;1,unknown;1,surprise;1,thought;1,policy;1,spelling;1,patrick;1,past;1,remains;1,pointed;1,lost;1,clued;1,puppeteer;1,control;1,handgun;1,robert;1,these;1,visibility;1,campaign;1,protections;1,past;1,manage;1,feathers;1,wilmington;1,difficult;1,elects;1,president;1,healy;1,submerge;1,documentary;1,feelings;1,candidacy;1,intellect;1,black;1,jacob;1,officer;1,ranch;1,how;1,main;1,trouble;1,ap;1,benjamin;1,needed;1,years;1,toward;1,messier;1,but;1,mouth;1,amassing;1,cars;1,off;1,goodell;1,take;1,night;1,southern;1,contact;1,happens;1,haven;1,influence;1,ally;1,figure;1,celebrating;1,takes;1,alina;1,climate;1,time;1,puzzles;1,erased;1,supporter;1,opposition;1,lose;1,rafiq;1,navigation;1,flag;1,some;1,ago;1,gift;1,kill;1,russia-ukraine;1,sheinbaum;1,undecided;1,plastic;1,press;1,ruling;1,tepetitán;1,partly;1,hostages;1,for;1,from;1,site;1,cease-fire;1,mail-in;1,before;1,español;1,war;1,attack;1,renewable;1,d-day;1,jeff;1,letting;1,germany;1,videos;1,polling;1,winnett;1,profile;1,dams;1,end;1,boggs;1,haiyun;1,formidable;1,nouf;1,academic;1,bought;1,current;1,whistle-blowers;1,window;1,low;1,matches;1,vote;1,made;1,revenge;1,expect;1,courthouse;1,reread;1,grillo;1,collective;1,gaza;1,parliament;1,division;1,henson;1,daily;1,oversee;1,becoming;1,seal;1,farewell;1,highways;1,jiang;1,one;1,shirt;1,cohen;1,asked;1,corrupt;1,ago;1,us;1,help;1,andrew;1,faces;1,hearing;1,electric;1,he;1,every;1,man;1,rises;1,paper;1,devastation;1,april;1,wu;1,vindictive;1,now;1,affordable;1,lost;1,totality;1,including;1,中文;1,comprehensive;1,flames;1,being;1,narrower;1,graff;1,index;1,suddenly;1,mumbai;1,b.j.p;1,insiders;1,chan;1,representative;1,employees;1,wire;1,work;1,unsettled;1,wordle;1,against;1,order;1,breaking;1,baring;1,wall;1,infectious;1,brexit;1,headed;1,paper;1,fallout;1,died;1,where;1,conviction;1,day;1,ad;1,metal;1,they;1,decay;1,between;1,former;1,cap;1,person;1,picture;1,thread;1,inspired;1,people;1,second;1,israel-hamas;1,stoking;1,margaret;1,chances;1,flies;1,celebrating;1,lawrence;1,probably;1,life;1,strike;1,know;1,finger;1,republican;1,studio;1,guru;1,night;1,newsletters;1,need;1,being;1,game;1,nytco;1,censored;1,campaign;1,swayed;1,ron;1,index;1,spotlight;1,destroying;1,ioan;1,transparency;1,bag;1,dead;1,dead;1,rose;1,floods;1,young;1,collins;1,college;1,would;1,keeps;1,contorted;1,top;1,tie;1,circle;1,written;1,tuesday;1,europe;1,close;1,dena;1,series;1,breaking;1,red;1,was;1,compound;1,crossword;1,behind;1,requiring;1,readers;1,people;1,map;1,opening;1,many;1,puzzle;1,make;1,border;1,largest;1,diversity;1,airstrikes;1,security;1,changes;1,going;1,way;1,tested;1,klein;1,display;1,welfare;1,mateos;1,early;1,a.i;1,voters;1,lawmakers;1,officials;1,weapons;1,get;1,rivera;1,too;1,ballots;1,idea;1,dry;1,e.c.i;1,stephens;1,author;1,charge;1,parties;1,commission;1,register;1,shunned;1,alliance;1,court;1,howard;1,calling;1,online;1,revolutionary;1,laying;1,trusts;1,various;1,order;1,hide;1,play;1,preferences;1,government;1,out;1,hinojosa;1,standing;1,daniel;1,with;1,stagnant;1,bath;1,isabel;1,arms;1,heat;1,mandates;1,defense;1,begins;1,statements;1,faculty;1,scope;1,some;1,openai;1,expected;1,storyful;1,source;1,disappears;1,solar;1,grilled;1,are;1,based;1,children;1,websites;1,suit;1,pursuit;1,surprisingly;1,caleb;1,kokotajlo;1,phillip;1,light;1,british;1,sweeping;1,lectern;1,tolga;1,memory;1,may;1,ezra;1,country;1,castigated;1,criminal;1,articles;1,stands;1,records;1,architect;1,we;1,personality;1,daunting;1,analysis;1,everyone;1,began;1,prior;1,mixed;1,most;1,supporters;1,drug;1,leaving;1,climate;1,visible;1,pandemic;1,current;1,discuss;1,circulated;1,scenario;1,hidden;1,yifan;1,contributed;1,tighter-than-expected;1,creature;1,keep;1,sign;1,amritsar;1,dominance;1,leader;1,additional;1,arnold;1,help;1,parts;1,stacks;1,set;1,expected;1,guess;1,by;1,thomas;1,debate;1,newsroom;1,south;1,n.d.a;1,jason;1,federal;1,background;1,break;1,convicted;1,explains;1,aid;1,had;1,hold;1,access;1,boom;1,answer;1,take;1,share;1,ruins;1,wind;1,ibrahim;1,felonies;1,advertise;1,documents;1,retire;1,come;1,came;1,country;1,watching;1,trends;1,glowing;1,stolen;1,up;1,letters;1,asylum;1,us;1,term;1,fence;1,galvanized;1,vaccine;1,ex-president;1,downplay;1,spent;1,large;1,stretches;1,proposal;1,lab;1,even;1,wave;1,points;1,archive;1,sheet;1,warn;1,shows;1,paws;1,helped;1,traced;1,scientists;1,audience;1,ramifications;1,companion;1,giving;1,fauci;1,keeps;1,surges;1,taking;1,beleaguered;1,claudia;1,just;1,blood;1,post;1,gail;1,months;1,history;1,says;1,life;1,group;1,one;1,only;1,garrett;1,none;1,coalition;1,economy;1,service;1,apartheid;1,wearing;1,become;1,bret;1,steals;1,theory;1,nearby;1,meet;1,smile;1,zuma;1,polling;1,what;1,nowhere;1,convicted;1,lifelong;1,turnaround;1,ayal;1,replica;1,muppets;1,begin;1,torso;1,companies;1,having;1,find;1,tarot;1,common;1,native;1,jacoby;1,up;1,fall;1,born;1,edition;1,information;1,gets;1,ensure;1,luntz;1,internet;1,law;1,mosque;1,whole;1,lying;1,curtice;1,true;1,sharp;1,anton;1,quiz;1,us;1,thumbs;1,funeral;1,led;1,aljowaysir;1,embassy;1,should;1,sale;1,britain;1,declared;1,creativity;1,cspan;1
```