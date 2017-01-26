package result 

import( 
    "ex"

    )
type StateCollector struct{
    sm StateManager 
}

func (sc *StateCollector)Consume(e ex.Event){
    sc.sm.UpdateState(e)
}
