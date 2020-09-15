import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateComponent } from './admin/create/create.component';
import { LoginComponent } from './admin/login/login.component';
import { RegisterComponent } from './admin/register/register.component';
import { MainviewComponent } from './mainview/mainview.component';
import { LandingComponent } from './landing/landing.component';
import { AboutComponent } from './about/about.component';


const adminRoutes: Routes = [
  {path:"create",component:CreateComponent}

];

const publicRoutes: Routes = [
  {path:"",component:LandingComponent},
  {path:"login",component:LoginComponent},
  {path:"register",component:RegisterComponent},
  {path:"about",component:AboutComponent}
]

const routes: Routes = [
  {path:"admin",children:adminRoutes,component:MainviewComponent},
  {path:"",children:publicRoutes,component:MainviewComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
