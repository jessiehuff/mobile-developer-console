const baseUrl = `/api`;

const fetchItems = url => {
    return fetch(url)
    .then(response => response.json())
    .then(result => result.items);
};

const dataService = {
    mobileClients: () => {
        return fetchItems(`${baseUrl}/mobileclients`);
    },
    serviceInstances: () => {
        return fetchItems(`${baseUrl}/serviceinstances`);
    },
    builds: () => {
        return fetchItems(`${baseUrl}/builds`);
    }
};

export default dataService;