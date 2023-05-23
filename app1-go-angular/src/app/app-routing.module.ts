import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateBookComponent } from './create-book/create-book.component';
import { BooklistComponent } from './booklist/booklist.component';
import { UpdateBookComponent } from './update-book/update-book.component';


const routes: Routes = [
  { path: '', redirectTo: 'view-book', pathMatch: 'full' },  
  { path: 'add-book', component: CreateBookComponent },
  { path: 'update-book/:data', component: UpdateBookComponent },
  {path:"**",component:BooklistComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
