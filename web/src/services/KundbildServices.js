import axios from 'axios';

// Fetch people, planets, and match data
export const fetchData = async (peopleId, planetsId) => {
    const peopleResponse = await axios.get(`http://localhost:8000/people/${peopleId}`);
    const planetsResponse = await axios.get(`http://localhost:8000/planets/${planetsId}`);
    const matchResponse = await axios.get(`http://localhost:8000/match/${peopleId}/${planetsId}`);

    return {
        people: peopleResponse.data,
        planets: planetsResponse.data,
        match: matchResponse.data.match
    };
};

// Save flagged customers
export const saveFlaggedCustomers = async (people, planets) => {
    await axios.post('http://localhost:8000/flagged_customers', { people, planets });
};

// Save linked customers
export const saveLinkedCustomers = async (people, planets) => {
    await axios.post('http://localhost:8000/linked_customers', { people, planets });
};
