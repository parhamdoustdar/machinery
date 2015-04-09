/*
 * Sending a task
 * --------------
 *
 * This simple example demonstrates sending a task to a Machinery worker.
 * In this case, a foobar task with foo and bar keyword arguments is sent.
 * It will be picked up and processed by one of worker processes.
 */

package main

import "github.com/RichardKnop/machinery/lib"

func main() {
	configMap := make(map[string]string)
	configMap["broker_url"] = "amqp://guest:guest@localhost:5672/"
	configMap["default_queue"] = "task_queue"

	app := lib.InitApp(configMap)
	
	// Send a test task
	name := "foobar"
	kwargs := make(map[string]interface{})
	kwargs["foo"] = "hello world"
	kwargs["bar"] = 123.4
	app.SendTask(name, kwargs)
}