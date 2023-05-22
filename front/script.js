// Get the table element
var table = document.getElementById("myTable");

// Specify the row and column indices you want to change the color for
var rowIndex = 1; // Index of the row (starting from 0)
var columnIndex = 2; // Index of the column (starting from 0)

// Get the cell in the specified row and column
var cell = table.rows[rowIndex].cells[columnIndex];

// Set the color of the cell
cell.style.backgroundColor = "#77e162";
cell.style.fontWeight = 'bold';

// Get the form element
var form = document.querySelector('form');

// Add an event listener to the form's submit event
form.addEventListener('submit', function(event) {
  event.preventDefault(); // Prevent the form from submitting normally

  // Get the selected values from the "frompoints" and "topoints" select elements
  var fromValue = document.getElementById('frompoints').value;
  var toValue = document.getElementById('topoints').value;

  // Create the URL with the selected values
  var url = 'http://localhost:8080/shortestpath?start=' + fromValue + '&end=' + toValue;

  // Send the request to the endpoint
  // You can use AJAX, Fetch API, or any other method to send the request
  // Here's an example using the Fetch API
  fetch(url)
  .then(function(response) {

    // Handle the response from the server
    return response.json(); // Parse response body as JSON
  })
  .then(function(data) {
    var table = document.getElementById('myTable');
    var rows = table.rows;
    for (var i = 0; i < rows.length; i++) {
      for (var j = 0; j < rows[i].cells.length; j++) {
        rows[i].cells[j].style.backgroundColor = ""; // Reset the background color
      }
    }
    // Access the parsed response data
    console.log(data);
    var fromPlace = table.rows[data.start+1].cells[0];
    fromPlace.style.backgroundColor = "#af9feb";
    var toPlace = table.rows[0].cells[data.end+1];
    toPlace.style.backgroundColor = "#af9feb";
    // You can do further processing with the data here
    
    for(var i=data.index.length-1; i > 0; i--){
        console.log("inside loop")
        const vert = data.index[i]+1
        const horiz = data.index[i-1]+1
        var cell = table.rows[vert].cells[horiz];
        
        cell.style.backgroundColor = "#82d992";
    }

    var distance  = document.getElementById('distance');
    distance.textContent = `${data.distance} km`
    var paths  = document.getElementById('paths');
    paths.textContent = `${data.path.join(" -> ")} `

  })
  .catch(function(error) {
    // Handle any errors that occur during the request
    console.log(error);
  });

});
