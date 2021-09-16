import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { UtilityService } from './utility.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'frontend';
  pageNumber = 1

  constructor(public http: HttpClient, public utilityService: UtilityService){

  }

  ngOnInit(): void {
    let that = this
    this.http.get('http://localhost:8080/api/getmoviecount/' ).subscribe(function(res: any){
      console.log("result of the movies count", res)
      let pages = Math.ceil(res/5)
      that.utilityService.moviesCount = pages;
  })

    this.http.get('http://localhost:8080/api/gettheatres/').subscribe(function(res){
      console.log("result of the theatre response", res)
      that.utilityService.theatres = res;
  })

    this.http.get('http://localhost:8080/api/movie/'+ this.pageNumber).subscribe(function(res){
      console.log("result of the movie response", res)
      that.utilityService.movies = res;
    })

    this.http.get('http://localhost:8080/api/getallmovies/').subscribe(function(res){
      console.log("result of the movie response", res)
      that.utilityService.getAllMovies = res;
    })

    // this.http.get('http://localhost:8080/api/getrecentmovies/').subscribe(function(res){
    //   console.log("result of the recent movies response", res)
    //   that.utilityService.recentMovies = res;
    //   // that.utilityService.movies = res;
    // })

    this.http.get('http://localhost:4200/api/checksession/').subscribe(function(res){
      console.log("result of the check session", res)
      that.utilityService.loggedInObj = res;
      // that.utilityService.movies = res;
    })

  }

  
}
