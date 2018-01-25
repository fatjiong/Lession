package main




type simpleCN struct{
	name string
	talk Talk
}

func NewSimpleCN(name string,talk Talk) ChatBoot{
	return &simpleCN{
		name:name,
		talk:talk,
	}
}

func (robot *simpleCN)Name()string  {
	return robot.name
}

func(robot *simpleCN)Begin()(string,error){

}

func (robot *simpleCN)Hello(userName string) string{
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
}

func (robot *simpleCN)ReportError(err error) string {

}

func (robot *simpleCN)End()error{

}