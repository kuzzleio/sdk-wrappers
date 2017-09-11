package main

//export kuzzle_wrapper_replay_queue
func kuzzle_wrapper_replay_queue() {
	KuzzleInstance.ReplayQueue()
}