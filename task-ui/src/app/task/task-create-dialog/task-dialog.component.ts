import { Component, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Task } from '../../services/task';
import { TaskService } from '../../services/task.service';
import {
  TASK_CREATED,
  TASK_ERROR,
  TASK_UPDATED
} from '../task-messages';
import { TaskDialogData } from './task-dialog-model';

@Component({
  selector: 'task-dialog',
  templateUrl: './task-dialog.component.html',
  styleUrl: './task-dialog.component.scss',
  standalone: false,
})
export class TaskDialogComponent {
  readonly dialogRef = inject(MatDialogRef<TaskDialogComponent>);
  readonly data = inject<TaskDialogData>(MAT_DIALOG_DATA);
  readonly snackBar = inject(MatSnackBar);

  task: Task = {} as Task;

  constructor(private taskService: TaskService) {}

  ngOnInit() {
    this.task = this.data.task || ({} as Task);
  }

  onConfirm() {
    if (this.data.type === 'create') {
      this.taskService.create(this.task).subscribe({
        next: () => {
          this.dialogRef.close(this.task);
          this.showSnackBar(TASK_CREATED);
        },
        error: (error) => {
          console.error('Error creating task', error);
          this.showSnackBar(TASK_ERROR);
        },
      });
      return;
    }

    this.taskService.update(this.task).subscribe({
      next: () => {
        this.dialogRef.close(this.task);
        this.showSnackBar(TASK_UPDATED);
      },
      error: (error) => {
        console.error('Error updating task', error);
        this.showSnackBar(TASK_ERROR);
      },
    });
  }

  showSnackBar(msg: string) {
    this.snackBar.open(msg, 'Close', {
      duration: 4000,
    });
  }
}
