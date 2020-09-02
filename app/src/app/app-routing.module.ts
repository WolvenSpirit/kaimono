import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateComponent } from './admin/create/create.component';


const adminRoutes: Routes = [
  {path:"",component:CreateComponent}
];

const routes: Routes = [
  {path:"admin",children:adminRoutes}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
