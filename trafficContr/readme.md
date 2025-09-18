

Classes:

Road
id 
name
trafficSignalId


SignalEnum: red , yellow , green

SetTrafficLight(light Light) // sets traffic light on a road

___________________
 
TrafficLight Class
id 
name
currentSignal:SignalEnum
RedDuration: int
GreenDuration:int
YellowDuration:int


durationofchange:int // seconds


ChangeSugnal(signal Sigal) // change signal and notifyobservors !!!!!!!!!!!!!!!!!!!!!
NotifyObservors()
_____________________________________

TrafficController // singleton, starts the trafic control process, manages roads
roads map[string] *Road
GetTrafficController() -- sigleton  to create the controller obj


AddRoad(road)  // add road to the controller class - map 

StartTrafficController (){
// for all the roads here , 
// for all trafic lights  --> call ChangeSignal()


 
}


HandleEmergency(){

}

_______________



TraffiCControllerDemo class

   trafficController :=GetController()
   r1:= NewRoad("r1","road1")

    r2:= NewRoad("r2","road2")
2
    s1:= NewLight("s1",10, 20,30) // duration of red , green , lights
    s2: NewLigh("s2",1,2,2)

    r1.SetATrafficLight(s1)

    trafficController.AddRoad(r1)
      trafficController.AddRoad(r2)

  go trafficController.Start()

  trafficController.HandleEmergency("R2")
  ______________________________________

