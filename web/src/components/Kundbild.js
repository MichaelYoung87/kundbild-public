import React, { useState } from 'react';
import { fetchData, saveFlaggedCustomers, saveLinkedCustomers } from '../services/KundbildServices';

function Kundbild() {
  const [peopleId, setPeopleId] = useState('');
  const [planetsId, setPlanetsId] = useState('');
  const [people, setPeople] = useState({});
  const [planets, setPlanets] = useState({});
  const [match, setMatch] = useState(null);
  const [statusMessage, setStatusMessage] = useState('');

  const handleFetchData = async () => {
    setStatusMessage(''); // Clear previous status messages

    // Only make API calls if peopleId and planetsId are valid
    if (peopleId >= 1 && peopleId <= 83 && planetsId >= 1 && planetsId <= 60) {
      try {
        const { people: fetchedPeople, planets: fetchedPlanets, match: fetchedMatch } = await fetchData(peopleId, planetsId);

        setPeople(fetchedPeople);
        setPlanets(fetchedPlanets);
        setMatch(fetchedMatch);
      } catch (err) {
        console.error(err);
        setStatusMessage('Error fetching data');
      }
    } else {
      setStatusMessage('Invalid people or planets ID');
    }
  };

  const handleSaveFlaggedCustomers = async () => {
    try {
      await saveFlaggedCustomers(people, planets);
      setStatusMessage('Flagged customers saved successfully');
    } catch (err) {
      console.error(err);
      setStatusMessage('Error saving flagged customers');
    }
  };

  const handleSaveLinkedCustomers = async () => {
    try {
      await saveLinkedCustomers(people, planets);
      setStatusMessage('Linked customers saved successfully');
    } catch (err) {
      console.error(err);
      setStatusMessage('Error saving linked customers');
    }
  };

  return (
    <div className="Kundbild">
      <input type="number" min="1" max="83" value={peopleId} onChange={(e) => {
        const val = e.target.value;
        if (val === '' || (!isNaN(val) && val >= 1 && val <= 83)) {
          setPeopleId(val);
          if (val === '') {
            setPeople({});
            setMatch(null);
            setStatusMessage('');
          }
        }
      }}
        placeholder="People URL Ending" />
      <input type="number" min="1" max="60" value={planetsId} onChange={(e) => {
        const val = e.target.value;
        if (val === '' || (!isNaN(val) && val >= 1 && val <= 60)) {
          setPlanetsId(val);
          if (val === '') {
            setPlanets({});
            setMatch(null);
            setStatusMessage('');
          }
        }
      }}
        placeholder="Planets URL Ending" />
      <button onClick={handleFetchData} disabled={!peopleId || !planetsId}>Fetch</button>
      <button disabled={!match} onClick={handleSaveFlaggedCustomers}>Save Flag</button>
      <button onClick={handleSaveLinkedCustomers} disabled={Object.keys(people).length === 0 || Object.keys(planets).length === 0}>Save Linked</button>

      {Object.keys(people).length > 0 && (
        <div>
          <h2>People Data:</h2>
          <pre>{JSON.stringify(people, null, 2)}</pre>
        </div>
      )}

      {Object.keys(planets).length > 0 && (
        <div>
          <h2>Planets Data:</h2>
          <pre>{JSON.stringify(planets, null, 2)}</pre>
        </div>
      )}

      {/* Display whether a match was found or not */}
      {match !== null && (
        <div>
          <h2>Match Result:</h2>
          <pre>{JSON.stringify(match, null, 2)}</pre>
        </div>
      )}
      {/* Display status message */}
      {statusMessage && (
        <div>
          <h2>Status:</h2>
          <p>{statusMessage}</p>
        </div>
      )}
    </div>
  );
}

export default Kundbild;