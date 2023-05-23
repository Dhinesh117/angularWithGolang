import { Component } from '@angular/core';
import { BookService } from '../book.service';
import { Book } from '../book';
@Component({
  selector: 'app-create-book',
  templateUrl: './create-book.component.html',
  styleUrls: ['./create-book.component.css']
})
export class CreateBookComponent {
title="Add Book"
constructor(private bookService:BookService) { }  
  
book : Book=new Book();  

onSubmit() {
  this.bookService.addBook(this.book).subscribe(
    data => {
      console.log(data);
    
    },
    error => console.log("Errpr throwing")
  );
  this.book = new Book();
}

   
}
