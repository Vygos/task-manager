import { HttpErrorResponse } from '@angular/common/http';
import { Component, HostListener, inject } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { PageEvent } from '@angular/material/paginator';
import { Page, Task } from '../services/task';
import { TaskService } from '../services/task.service';
import { TaskDialogComponent } from './task-create-dialog/task-dialog.component';
import { MatSnackBar } from '@angular/material/snack-bar';
import { TASK_DELETED, TASK_ERROR } from './task-messages';

@Component({
  selector: 'task-list',
  templateUrl: './task-list.component.html',
  styleUrl: './task-list.component.scss',
  standalone: false,
})
export class TaskListComponent {
  DEFAULT_PAGE_SIZE = 5;

  page: Page<Task> = { data: [], total_elements: 0, page: 0, size: 3 };
  columns = ['title', 'status', 'created_at', 'updated_at', 'actions'];
  isMobile: boolean = window.innerWidth < 768;

  readonly dialog = inject(MatDialog);
  readonly snackBar = inject(MatSnackBar);

  constructor(private readonly taskService: TaskService) {}

  ngOnInit() {
    this.taskService
      .getAll({
        size: this.DEFAULT_PAGE_SIZE,
      } as Page<Task>)
      .subscribe({
        next: (taskPage) => {
          this.page = taskPage;
        },
        error: (error: HttpErrorResponse) => {
          console.error('Error fetching tasks', error);
          this.showSnackBar(TASK_ERROR);
        }
      });
  }

  @HostListener('window:resize', ['$event'])
  onResize() {
    this.isMobile = window.innerWidth <= 768;
    if (this.isMobile) {
      this.columns = ['title', 'status', 'actions'];
    } else {
      this.columns = ['title', 'status', 'created_at', 'updated_at', 'actions'];
    }
  }

  onPageChange(pageEvent: PageEvent) {
    this.taskService
      .getAll({
        page: pageEvent.pageIndex + 1,
        size: pageEvent.pageSize,
      } as Page<Task>)
      .subscribe({
        next: (taskPage) => {
          this.page = taskPage;
        },
        error: (error: HttpErrorResponse) => {
          console.error('Error fetching tasks', error);
          this.showSnackBar(TASK_ERROR);
        }
      });
  }

  onAddTask() {
    const dialogRef = this.dialog.open(TaskDialogComponent, {
      data: {
        type: 'create',
        title: 'Create Task',
      },
    });

    dialogRef.afterClosed().subscribe(() => {
      this.taskService
        .getAll({
          size: this.DEFAULT_PAGE_SIZE,
          page: this.page.page,
        } as Page<Task>)
        .subscribe({
          next: (taskPage) => {
            this.page = taskPage;
            console.log('Task updated');
          },
          error: (error: HttpErrorResponse) => {
            console.error('Error Fetching task page task', error);
          },
        });
    });
  }

  onEditTask(task: Task) {
    const dialogRef = this.dialog.open(TaskDialogComponent, {
      data: { title: 'Edit task', type: 'edit', task: {...task} },
    });

    dialogRef.afterClosed().subscribe(() => {
      this.taskService
        .getAll({
          size: this.DEFAULT_PAGE_SIZE,
          page: this.page.page,
        } as Page<Task>)
        .subscribe({
          next: (taskPage) => {
            this.page = taskPage;
          },
          error: (error: HttpErrorResponse) => {
            console.error('Error deleting task', error);
            this.showSnackBar(TASK_ERROR);
          },
        });
    });
  }

  onDeleteTask(task: Task) {
    this.taskService.delete(task.id).subscribe({
      next: () => {
        this.taskService
          .getAll({
            size: this.DEFAULT_PAGE_SIZE,
            page: this.page.page,
          } as Page<Task>)
          .subscribe((taskPage) => {
            this.page = taskPage;
          });
        this.showSnackBar(TASK_DELETED);
      },
      error: (error: HttpErrorResponse) => {
        console.error('Error deleting task', error);
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
