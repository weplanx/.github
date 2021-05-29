import { NgModule } from '@angular/core';
import { AppShareModule } from '@share';
import { RouterModule, Routes } from '@angular/router';
import { ProjectsIndexComponent } from './projects-index/projects-index.component';
import { ProjectsPageComponent } from './projects-page/projects-page.component';
import { ProjectsArchiveComponent } from './projects-archive/projects-archive.component';

const routes: Routes = [
  {
    path: '',
    component: ProjectsIndexComponent
  },
  {
    path: 'archive',
    component: ProjectsArchiveComponent
  },
  {
    path: 'key/:key',
    component: ProjectsPageComponent
  }
];

@NgModule({
  imports: [
    AppShareModule,
    RouterModule.forChild(routes)
  ],
  declarations: [
    ProjectsIndexComponent,
    ProjectsArchiveComponent,
    ProjectsPageComponent
  ],
  exports: [
    ProjectsIndexComponent,
    ProjectsArchiveComponent,
    ProjectsPageComponent
  ]
})
export class ProjectsModule {
}
