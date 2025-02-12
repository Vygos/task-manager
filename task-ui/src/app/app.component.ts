import { Component } from "@angular/core";
import { TaskModule } from "./task/task.module";


@Component({
    selector: "app-root",
    imports: [TaskModule],
    template: "<task-list></task-list>",
})
export class AppComponent {

}