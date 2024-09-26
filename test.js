// TESTS IF get_platforms WORKS
// fetch('http://127.0.0.1:3000/api/get-platforms', {
// 	method: 'GET',
// 	headers: {
// 		'Content-Type': 'application/json'
// 	}
// })
// 	.then(response => response.json())
// 	.then(data => console.log(data))
// 	.catch(error => console.error(error));


// TESTS IF get_platform_by_name WORKS
// fetch('http://127.0.0.1:3000/api/get-platform-by-name', {
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/json'
// 	},
// 	body: JSON.stringify({ name: 'notaplatform' })
// })
// 	.then(response => response.json())
// 	.then(data => console.log(data))
// 	.catch(error => console.error(error));


// TESTS IF get_platform_by_id WORKS
// fetch('http://127.0.0.1:3000/api/get-platform-by-id', {
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/json'
// 	},
// 	body: JSON.stringify({ id: '66eb0fae96ad1476e9e20c55' })
// })
// 	.then(response => response.json())
// 	.then(data => console.log(data))
// 	.catch(error => console.error(error));


// TESTS IF add_platform WORKS
// fetch('http://127.0.0.1:3000/api/add-platform', {
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/json'
// 	},
// 	body: JSON.stringify({
// 		name: 'notaplatform',
// 		manufacturer: 'notamanufacturer',
// 	})
// })
// 	.then(response => response.json())
// 	.then(data => console.log(data))
// 	.catch(error => console.error(error));


// TESTS IF update_platform WORKS
// fetch('http://127.0.0.1:3000/api/update-platform', {
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/json'
// 	},
// 	body: JSON.stringify({
// 		id: '66f18b44cd9cfa6aa59f5584',
// 		name: 'Saturn',
// 		manufacturer: 'Sega'

// 	})
// }
// 	.then(response => response.json())
// 	.then(data => console.log(data))
// 	.catch(error => console.error(error));

// TESTS IF delete_platform WORKS
fetch('http://127.0.0.1:3000/api/delete-platform', {
	method: 'POST',
	headers: {
		'Content-Type': 'application/json'
	},
	body: JSON.stringify({
		id: '66f528e30a6ed919a18f24f6'
	})
})
	.then(response => response.json())
	.then(data => console.log(data))
	.catch(error => { console.log(error) })