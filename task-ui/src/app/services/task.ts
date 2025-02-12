export interface Page<T> {
  data: T[];
  total_elements: number;
  page: number;
  size: number;
}

export interface Task {
  id: string;
  title: string;
  status: string;
  created_at: Date;
  updated_at: Date;
}
