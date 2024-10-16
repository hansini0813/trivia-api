// Add click event to fact to show answer
// Get the editForm element and the current fact ID to edit
const editForm = document.querySelector('#form-update-fact')
const factToEdit = editForm && editForm.dataset.factid

// Add an event listener to listen for the form submit
editForm && editForm.addEventListener('submit', (event) => {
	// Prevent the default behaviour of the form element
    event.preventDefault()

	// Convert form data into a JavaScript object
    const formData = Object.fromEntries(new FormData(editForm));

    return fetch(`/fact/${factToEdit}`, {
		// Use the PATCH method
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
		// Convert the form's Object data into JSON
        body: JSON.stringify(formData),
    })
    .then(() => document.location.href=`/fact/${factToEdit}`)// Redirect to show
})
