import { Component, OnInit } from '@angular/core';
import { BookService } from '../book.service';
import { Book } from '../book';
import { Route, Router } from '@angular/router';
@Component({
  selector: 'app-booklist',
  templateUrl: './booklist.component.html',
  styleUrls: ['./booklist.component.css']
})
export class BooklistComponent implements OnInit{

 
  book : Book=new Book();  

  constructor(private bookService:BookService, private route : Router ) { }  
  
  books: any[] = [];  
  ngOnInit() {  
    this.bookService.getBookList().subscribe(data =>{  
      this.books =data;  
  })
}
deleteMessage=false;  

deleteBook(id: string) {  
  // let id1= Number(id)
  // this.bookService.deleteBook(v).subscribe(data => {
  //   console.log(data);
  // });
  this.bookService.deleteBook(id)  
    .subscribe(  
      data => {  
        console.log(data);  
        this.deleteMessage=true;  
        this.bookService.getBookList().subscribe(data =>{  
          this.books =data  
          })  
      },  
      error => console.log(error));  
}  

updateBook(book: Book){
  console.log(">>>>>>>>>>>>>>>>>>>")
  let data=JSON.stringify(book)
  this.route.navigate(['/update-book/'+data])
}
}