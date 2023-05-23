import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http'
import { Book } from './book';
@Injectable({
  providedIn: 'root'
})
export class BookService {

  private baseUrl='http://localhost:8000/books';
  private baseUrl1='http://localhost:8000/book';

  constructor(private httpClient: HttpClient) { }

  getBookList(): Observable<any>{
    return this.httpClient.get<string>(this.baseUrl);
  }

  addBook(book: Book): Observable<any> {  
    let id1 = Number(book.id);
    let price1=Number(book.price);

    // Create a new book object with the integer id
    const newBook = {
      id: id1,
      author:book.author,
      title:book.title,
      price:price1
    };

    return this.httpClient.post(`${this.baseUrl1}`, newBook);  
  }  
  deleteBook(id1: string): Observable<any> {  
 
    return this.httpClient.delete(`${this.baseUrl}/${id1}`, { responseType: 'text' });  
  }  
  
  getBook(id :number): Observable<any> {  
    return this.httpClient.get(`${this.baseUrl}/${id}`, { responseType: 'json'});  
  }  
  updateBook(book: Book):Observable<any>{
    let id1 = Number(book.id);
    let price1=Number(book.price);
    const newBook = {
      id: id1,
      title:book.title,
      author:book.author,
      price:price1
    };
  
    return this.httpClient.put(`${this.baseUrl}/${id1}`, newBook);  
  }  
  }

