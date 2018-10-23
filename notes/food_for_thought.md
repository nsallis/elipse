# Food for thought

## hooking up a node with multiple children

* we should create channels for each node, and then look at the child component, and use its input channel as the parent's output channel. That way we can easily accomidate multiple children down the road

## ordered nodes

* will need a way to keep order of splits even when we split a document multiple times. Will start with only unordered nodes.

## need logging

## need to be able to save state when we exit

* either save the state of each node so we can restore it later, or at least ask to confirm exiting unsaved nodes
