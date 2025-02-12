import { Task } from "../../services/task";

export interface TaskDialogData {
    title: string;
    type: 'create' | 'edit';
    task?: Task;
}