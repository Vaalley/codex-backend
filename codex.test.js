// Test for checking if the endpoint get_platforms works and returns an array of platforms
test("get_platforms", async () => {
	const response = await fetch('http://127.0.0.1:3000/api/get-platforms', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	})
		.then(response => response.json());

	expect(response).toBeInstanceOf(Array);
	response.forEach(platform => {
		expect(typeof platform.name).toBe('string')
		expect(typeof platform.manufacturer).toBe('string')
		expect(typeof platform.ID).toBe('string')
		expect(typeof platform.type).toBe('string')
	});
});

// Test for checking if the endpoint get_platform_by_name works and returns a single platform object
test("get_platform_by_name", async () => {
	const response = await fetch('http://127.0.0.1:3000/api/get-platform-by-name', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ name: 'Dreamcast' })
	})
		.then(response => response.json());

	expect(response).toBeInstanceOf(Object);
	expect(response).toHaveProperty('name');
	expect(response).toHaveProperty('manufacturer');
	expect(response).toHaveProperty('ID');
	expect(response).toHaveProperty('type');
})

// Test for checking if the endpoint get_platform_by_id works and returns a single platform object
test("get_platform_by_id", async () => {
	const response = await fetch('http://127.0.0.1:3000/api/get-platform-by-id', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ id: '66eb0fae96ad1476e9e20c55' })
	})
		.then(response => response.json());

	expect(response).toBeInstanceOf(Object);
	expect(response).toHaveProperty('name');
	expect(response).toHaveProperty('manufacturer');
	expect(response).toHaveProperty('ID');
	expect(response).toHaveProperty('type');
})

// Test for checking if the endpoints add_platform, update_platform, and delete_platform work
test('add, update, and delete platform', async () => {
	// Add platform
	const addResponse = await fetch('http://127.0.0.1:3000/api/add-platform', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			name: 'notaplatform',
			manufacturer: 'notamanufacturer',
		}),
	});
	const addedPlatform = await addResponse.json();
	expect(addedPlatform).toHaveProperty('ID');

	// Update platform
	const updateResponse = await fetch('http://127.0.0.1:3000/api/update-platform', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			id: `${addedPlatform.ID}`,
			name: 'updatedplatform',
			manufacturer: 'updatedmanufacturer',
		}),
	});
	const updatedPlatform = await updateResponse.json();
	expect(updatedPlatform).toHaveProperty('ID');
	expect(updatedPlatform.ID).toBe(addedPlatform.ID);

	// Delete platform
	const deleteResponse = await fetch('http://127.0.0.1:3000/api/delete-platform', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			id: `${addedPlatform.ID}`,
		}),
	});
	const deletedPlatform = await deleteResponse.json();
	expect(deletedPlatform.message).toBe('Platform deleted successfully');
})
