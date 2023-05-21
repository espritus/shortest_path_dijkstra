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