import { TaskProps, TaskResponseProps } from "@/types/Task";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosPromise } from "axios";

const fetchData = async (): AxiosPromise<TaskResponseProps> =>{
  const response = await axios.get("http://localhost:8083/api/v1/task/list")
  return response
}

export function useTaskList(){
  const query = useQuery({
    queryFn: fetchData,
    queryKey: ['task-data']
  })

  const task:TaskProps[]| any = query.data?.data

  return {
    ...query,
    data: task
  }
}

