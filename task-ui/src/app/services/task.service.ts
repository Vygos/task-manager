import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Page, Task } from './task';
import { environments } from '../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class TaskService {
  private apiUrl = environments.apiUrl;

  constructor(private readonly http: HttpClient) {}

  getAll(page?: Page<Task>): Observable<Page<Task>> {
    return this.http.get<Page<Task>>(`${this.apiUrl}/tasks`, {params: {
      page: page?.page || 1,
      size: page?.size || 3,
    }});
  }

  create(task: Task): Observable<Task> {
    return this.http.post<Task>(`${this.apiUrl}/tasks`, {
      title: task.title,
      status: task.status,
    });
  }

  update(task: Task): Observable<Task> {
    return this.http.patch<Task>(`${this.apiUrl}/tasks/${task.id}`, {
      title: task.title,
      status: task.status,
    });
  }

  delete(id: string): Observable<Task> {
    return this.http.delete<Task>(`${this.apiUrl}/tasks/${id}`);
  }
}
