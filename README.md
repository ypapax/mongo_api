this is my solution of an interview task:
# The task

1) create a go service that reads from a YAML config file:

apiKeys:
- a
- b
- couldBeOneOrMore
- reallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylongreallylong
listenPort: 80
mongoDb:
  connectionString: “mongodb://localhost:27017”
  targetCollection: "auditlogs"

2) app shall start an multi-threaded http server on the configured port

3) the server shall accept a REST call (POST to /auditlog) that requires a JSON structure, i.e.:

{ 
 	"apiKey": "7GXv.^7PWJkHW}4:(UW?jWA5bDV+=q;JV&z+'&#)Q6,XR7VpW`nTL/t,*9@?^?Fa>TC=b2gX%/&~xg\Ej@/)5J~r)m9RhrMfvt2g$qfx89>%e4$?P$_(a5B7RGdk2ksEj@RRtwP;P[%YF.X5'SQmc<H7XH^JD>89TNh-nN*HY[^bm+Yq.=~[yzD)n?Y,dbe3uz&Bb;q"~gkmH(b3s8#a\Dy.jDM`pER?j52eS[YdYAz2/ST\yAL\"dys3Spa'rUU]r%ZB)f8q8Gs!B!zrZCH+bRs;4NnR(gF.rW]h4.LrrF/tYB?y4U<FGLK6cA}(6DZ;AWf]%Q2w6jW~(){mz[sBD-<YSH#Eh+u"C.W#p?'{SZ'4CkCAV*TW8~]4p4?p.nfLmSjsYC.WF{6Z_zS+vQE9yEqr?@`^VZ:QWC?A*^7')k+y?@J:</YF8MTzg>x]9dM~@2#N+s/'}z?:m=TmdH6kf]K%yY_uE({]UYY4VBk6[2`s)9?vu!cLj5uaR?2jFZ/",
    "url" : "/order/create", 
    "companyId" : "5ba0d465855a51722785f967", 
	"object" : "something-as-this-field-is-optional",
    "action" : "OrderCreate", 
    "data" : "[{\"_id\":\"5cf97d9be3e361e0e4aaff97\",\"index\":0,\"guid\":\"2f2ef81a-4ad4-4ac0-843a-33b1f53812e7\",\"isActive\":false,\"balance\":\"$2,827.22\",\"picture\":\"http://placehold.it/32x32\",\"age\":31,\"eyeColor\":\"green\",\"name\":\"Leblanc Palmer\",\"gender\":\"male\",\"company\":\"COMVENE\",\"email\":\"leblancpalmer@comvene.com\",\"phone\":\"+1 (902) 458-3168\",\"address\":\"647 Losee Terrace, Blodgett, Pennsylvania, 6956\",\"about\":\"Ullamco Lorem elit labore culpa cupidatat duis irure labore adipisicing. Mollit elit irure magna Lorem in. Laborum reprehenderit aute ut incididunt. Labore adipisicing esse nisi excepteur culpa. Ut ex esse nulla officia sit ex in. Irure qui non nulla Lorem cupidatat incididunt.\r\n\",\"registered\":\"2017-12-14T11:12:27 -01:00\",\"latitude\":27.555353,\"longitude\":-128.787282,\"tags\":[\"officia\",\"velit\",\"ullamco\",\"do\",\"id\",\"adipisicing\",\"nisi\"],\"friends\":[{\"id\":0,\"name\":\"Raymond Olsen\"},{\"id\":1,\"name\":\"Rachel Massey\"},{\"id\":2,\"name\":\"Zelma Kent\"}],\"greeting\":\"Hello, Leblanc Palmer! You have 2 unread messages.\",\"favoriteFruit\":\"strawberry\"},{\"_id\":\"5cf97d9b53b999cbc548f990\",\"index\":1,\"guid\":\"0b90eb61-1341-4399-ade1-d1d27009321e\",\"isActive\":true,\"balance\":\"$2,361.61\",\"picture\":\"http://placehold.it/32x32\",\"age\":28,\"eyeColor\":\"green\",\"name\":\"April Patton\",\"gender\":\"female\",\"company\":\"SLOGANAUT\",\"email\":\"aprilpatton@sloganaut.com\",\"phone\":\"+1 (902) 471-2069\",\"address\":\"463 Calyer Street, Winchester, Louisiana, 9616\",\"about\":\"Velit irure cillum veniam ad excepteur aute duis ipsum Lorem commodo consequat. Deserunt eu sint proident nostrud velit eiusmod. Ipsum ullamco eu aliqua labore non aliqua pariatur ullamco dolor tempor enim elit cillum.\r\n\",\"registered\":\"2015-05-25T12:17:57 -02:00\",\"latitude\":14.027941,\"longitude\":-130.498889,\"tags\":[\"aute\",\"excepteur\",\"aute\",\"ullamco\",\"sint\",\"laborum\",\"velit\"],\"friends\":[{\"id\":0,\"name\":\"Luella Rose\"},{\"id\":1,\"name\":\"Ruth Mcknight\"},{\"id\":2,\"name\":\"Valenzuela Carr\"}],\"greeting\":\"Hello, April Patton! You have 4 unread messages.\",\"favoriteFruit\":\"banana\"},{\"_id\":\"5cf97d9bdfbb0f0a46aa1d1c\",\"index\":2,\"guid\":\"dd216c90-c241-41dc-b068-24d1e05e9b09\",\"isActive\":false,\"balance\":\"$2,095.73\",\"picture\":\"http://placehold.it/32x32\",\"age\":21,\"eyeColor\":\"brown\",\"name\":\"Alice Morton\",\"gender\":\"female\",\"company\":\"LETPRO\",\"email\":\"alicemorton@letpro.com\",\"phone\":\"+1 (930) 437-3048\",\"address\":\"236 Emerson Place, Forestburg, Kentucky, 3585\",\"about\":\"Ex laboris cupidatat adipisicing magna amet voluptate Lorem. Esse irure proident labore tempor cillum duis. Veniam in ut fugiat quis consectetur. Culpa qui incididunt ex ullamco incididunt deserunt nulla voluptate eu. Proident ea exercitation anim est est. Labore et nisi ullamco pariatur aute exercitation sint elit ullamco. Magna occaecat occaecat adipisicing commodo ex id Lorem exercitation pariatur eu exercitation labore consectetur et.\r\n\",\"registered\":\"2017-11-15T04:24:41 -01:00\",\"latitude\":6.12763,\"longitude\":-26.461493,\"tags\":[\"consectetur\",\"dolor\",\"incididunt\",\"elit\",\"ipsum\",\"ullamco\",\"adipisicing\"],\"friends\":[{\"id\":0,\"name\":\"Baker Manning\"},{\"id\":1,\"name\":\"Cheri Ortega\"},{\"id\":2,\"name\":\"Wiggins Neal\"}],\"greeting\":\"Hello, Alice Morton! You have 8 unread messages.\",\"favoriteFruit\":\"strawberry\"},{\"_id\":\"5cf97d9b8f81dc88daa3f6c0\",\"index\":3,\"guid\":\"b44a38fb-e242-486b-84c2-2c751b3ffc9c\",\"isActive\":true,\"balance\":\"$1,394.76\",\"picture\":\"http://placehold.it/32x32\",\"age\":40,\"eyeColor\":\"green\",\"name\":\"Lilia Best\",\"gender\":\"female\",\"company\":\"EXOSIS\",\"email\":\"liliabest@exosis.com\",\"phone\":\"+1 (998) 550-2426\",\"address\":\"666 Hanover Place, Johnsonburg, Marshall Islands, 8012\",\"about\":\"Excepteur duis nulla Lorem quis aliqua deserunt duis nostrud duis sit ex commodo amet cillum. Ea occaecat Lorem culpa aliquip sit commodo exercitation. Ipsum non nisi non in do. Non eiusmod consequat reprehenderit anim ullamco veniam aute eu est incididunt ad.\r\n\",\"registered\":\"2014-10-21T07:03:45 -02:00\",\"latitude\":83.845016,\"longitude\":-53.898336,\"tags\":[\"exercitation\",\"reprehenderit\",\"pariatur\",\"id\",\"excepteur\",\"et\",\"culpa\"],\"friends\":[{\"id\":0,\"name\":\"Rosalind Hurst\"},{\"id\":1,\"name\":\"Knight Tran\"},{\"id\":2,\"name\":\"Nona Tanner\"}],\"greeting\":\"Hello, Lilia Best! You have 9 unread messages.\",\"favoriteFruit\":\"banana\"},{\"_id\":\"5cf97d9bcfbf329736384157\",\"index\":4,\"guid\":\"4ca50699-a020-4226-a7fc-abaca44aa00e\",\"isActive\":false,\"balance\":\"$3,664.83\",\"picture\":\"http://placehold.it/32x32\",\"age\":22,\"eyeColor\":\"blue\",\"name\":\"Eunice Dickson\",\"gender\":\"female\",\"company\":\"ENDICIL\",\"email\":\"eunicedickson@endicil.com\",\"phone\":\"+1 (869) 439-3905\",\"address\":\"104 Matthews Place, Orviston, New Hampshire, 3801\",\"about\":\"Laborum sunt minim consectetur in excepteur deserunt. Anim dolor pariatur Lorem nulla nulla magna consequat do aliquip ullamco esse ea velit consequat. Laborum nisi aliquip magna irure aute aliqua aute et anim duis ex.\r\n\",\"registered\":\"2018-10-08T08:27:28 -02:00\",\"latitude\":-58.016872,\"longitude\":36.509332,\"tags\":[\"enim\",\"enim\",\"nostrud\",\"dolor\",\"veniam\",\"laborum\",\"ex\"],\"friends\":[{\"id\":0,\"name\":\"Jordan Strickland\"},{\"id\":1,\"name\":\"Kayla Pitts\"},{\"id\":2,\"name\":\"Hawkins Webb\"}],\"greeting\":\"Hello, Eunice Dickson! You have 5 unread messages.\",\"favoriteFruit\":\"strawberry\"}]", 
    "createdBy" : "5ba0d465855a51722785f975", 
    "createdByFullName" : "Zorro", 
    "createdOn" : ISODate("2019-06-01T06:10:39.400+0000"), 
    "sourceSystem" : "WEB", 
    "systemVersion" : 1975.0, 
    "uuid" : "5cf216d3adfe582cdb06a97f", 
    "auditDirection" : "incoming", 
    "requestId" : NumberInt(832788)
}

4) check if apiKey is known
-> if not found within the values from the config files, respond with 401 unauthorized error and ignore
-> if known, persist record into the specified collection (everything received except apiKey)

# Solution

## Building
./commands.sh build

## Running
./commands.sh run

## Testing
### success response:
./commands.sh test

### failed 401:
./commands.sh test2
