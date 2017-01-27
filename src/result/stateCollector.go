package result 

import( 
    "ex"

    )
type StateCollector struct{
    Sm StateManager 
}

func (sc *StateCollector)Consume(e ex.Event){
    sc.Sm.UpdateState(e)
}
