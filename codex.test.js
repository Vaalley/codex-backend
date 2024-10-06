// Test for checking if the endpoint get_platforms works and returns an array of platforms
test("get_platforms", async () => {
	const response = await codexFetch('get-platforms', { method: 'GET' })

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
	const response = await codexFetch('get-platform-by-name', { body: { name: 'Dreamcast' } })

	expect(response).toBeInstanceOf(Object);
	expect(response).toHaveProperty('name');
	expect(response).toHaveProperty('manufacturer');
	expect(response).toHaveProperty('ID');
	expect(response).toHaveProperty('type');
})

// Test for checking if the endpoint get_platform_by_id works and returns a single platform object
test("get_platform_by_id", async () => {
	const response = await codexFetch('get-platform-by-id', {
		body: { id: '66eb0fae96ad1476e9e20c55' }
	})

	expect(response).toBeInstanceOf(Object);
	expect(response).toHaveProperty('name');
	expect(response).toHaveProperty('manufacturer');
	expect(response).toHaveProperty('ID');
	expect(response).toHaveProperty('type');
})

// Test for checking if the endpoints add_platform, update_platform, and delete_platform work
test('add, update, and delete platform', async () => {
	// Add platform
	const addResponse = await codexFetch('add-platform', {
		body: {
			name: 'notapllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllatform',
			manufacturer: 'notamanufacturer',
		}
	})
	if (addResponse.error) {
		expect(addResponse).not.toHaveProperty('error');
		return
	}
	expect(addResponse).toHaveProperty('ID');

	// Update platform
	const updateResponse = await codexFetch('update-platform', {
		body: {
			id: `${addResponse.ID}`,
			name: 'updatedplatform',
			manufacturer: 'updatedmanufacturer',
		}
	})
	expect(updateResponse).toHaveProperty('ID');
	expect(updateResponse.ID).toBe(addResponse.ID);

	// Delete platform
	const deleteResponse = await codexFetch('delete-platform', {
		body: {
			id: `${updateResponse.ID}`,
		}
	})
	expect(deleteResponse.message).toBe('Platform deleted successfully');
})

function codexFetch(endpoint, params = {}) {
	return fetch(`http://127.0.0.1:3000/api/${endpoint}`, {
		method: 'POST',
		...params,
		body: JSON.stringify(params.body),
		headers: {
			'Content-Type': 'application/json',
			...params.headers
		},
	}).then(response => response.json());
}