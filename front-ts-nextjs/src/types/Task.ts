export type TaskProps ={
  ID: number,
  name: string,
  isDone: boolean
}

export interface TaskResponseProps {
  data: TaskProps[] |null
}