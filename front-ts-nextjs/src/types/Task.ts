export type TaskProps ={
  ID: number,
  name: string,
  isDone: boolean
}

export type TaskResponseProps ={
  data:TaskProps[]
}