import { Component } from '@angular/core';
import { BookService } from '../book.service';
import { Book } from '../book';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-update-book',
  templateUrl: './update-book.component.html',
  styleUrls: ['./update-book.component.css']
})
export class UpdateBookComponent {
  constructor(private bookService:BookService,private route:ActivatedRoute) { }  
book1:any
parseDat:any
// book1 : Book=new Book();  
ngOnInit() {
  let data=this.route.snapshot.paramMap.get('data')
  this.parseDat=data
  this.book1=JSON.parse(this.parseDat)
  console.log("*******************",this.book1)
}

onSubmit() {
  this.bookService.updateBook(this.book1).subscribe(
   
    data => {
      data=this.book1;
      console.log("+++++++++++++++++++++++++",data);
    
    },
    error => console.log("Errpr throwing")
  );
  // this.book1 = new Book();
}
}