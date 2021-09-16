import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { UtilityService } from '../utility.service';
import { DomSanitizer, SafeResourceUrl, SafeUrl} from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  movieName: any
  pageNumber = 1
  constructor(public utilityService: UtilityService, public http: HttpClient, private sanitizer : DomSanitizer) {
    
   }

  ngOnInit(): void {
    
  }

  nextPage() {
    let that = this;
    this.pageNumber = this.pageNumber + 1
    this.http.get('http://localhost:8080/api/movie/' + this.pageNumber).subscribe(function(res){
      console.log("result of the movie response", res)
      that.utilityService.movies = res;
    })
  }

  previousPage() {
    let that = this
    this.pageNumber = this.pageNumber -1
    this.http.get('http://localhost:8080/api/movie/' + this.pageNumber).subscribe(function(res){
      console.log("result of the movie response", res)
      that.utilityService.movies = res;
    })
  }

}
