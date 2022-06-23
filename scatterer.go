package drifter

// A Scatterer knows how to add Traces to a Sim, distributed in various ways (random, regular intervals).
type Scatterer interface{}

// NOTES
//
// Needs to know when to run or be able to make that decision based on info passed from the Sim. What do we need to pass?
//
// Who decides when a trace added by a scatterer should become inactive?
