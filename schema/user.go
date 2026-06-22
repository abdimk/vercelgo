package schema

import "time"


type Emotions struct{
	Name string `json:"name"` // emotion name
	Intensity float64 `json:"Intensity"` // 
	Valence     float64 `json:"valence"` 
    Duration    time.Duration `json:"Duration"`
    Trigger     string `json:"trigger"`
    StartTime   time.Time `json:"start_time"`
    Target      string `json:"target"`
    PhysicalEffects []string `json:"physical_effects"`
    Thoughts    []string `json:"toughts"`
    Expression  string `json:"expression"`
	
}