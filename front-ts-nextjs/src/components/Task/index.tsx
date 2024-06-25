'use server'

import { useTaskList } from "@/hooks/useTaskList";
import { TaskProps } from "@/types/Task";


const Task = () => {
  const {data } = useTaskList()
  console.log(data)
  return (
    <>
    {data?.data.data.map(task=>{
      <>
        <h1>{task.name}</h1>
      </>
    })}
    </>
  )
}

export default Task