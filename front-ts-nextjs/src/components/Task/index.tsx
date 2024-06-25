"use client"
import { useTaskList } from "@/hooks/useTaskList";
import { TaskProps, TaskResponseProps } from "@/types/Task";


const Task = () => {
  const { data } = useTaskList()
  console.log("dataaaaaaaaaaaaa:", data)

  const handlerCheck = ()=>{
    const isChecked = data?.data.map((task:TaskProps) =>{task.isDone})
    return isChecked
  }

  return (
    <>
      <h1>asdasd</h1>
      {data?.map((task: TaskProps) => (
        <div key={task.ID}>
          <h1 key={task.ID}>{task.name}</h1>
          <input key={task.ID} type="checkbox" checked={task.isDone}/>
        </div>
      ))}
    </>
  )
}

export default Task