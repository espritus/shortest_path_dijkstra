// Make an HTTP request to retrieve the JSON data
fetch('localhost:9090/shortestpath')
    .then(response => response.json())
    .then(data => {
        // Process the received JSON data here
        console.log(data);
        // Access specific properties of the JSON object
        var start = data.start;
        var end = data.end;
        var path = data.path;
        var distance = data.distance;
        var index = data.index;

        var table = document.getElementById("myTable");

        // Specify the row and column indices you want to change the color for
        var rowIndex = 2; // Index of the row (starting from 0)
        var columnIndex = 6; // Index of the column (starting from 0)

        // Get the cell in the specified row and column
        var cell = table.rows[rowIndex].cells[columnIndex];

        // Set the color of the cell
        cell.style.backgroundColor = "#77e162";
        cell.style.fontWeight = 'bold';
    })
    .catch(error => {
        // Handle any errors that occurred during the request
        console.log('Error:', error);
    });
