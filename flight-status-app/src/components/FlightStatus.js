// src/components/FlightStatus.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const FlightStatus = () => {
  const [flights, setFlights] = useState([]);

  useEffect(() => {
    axios.get('/api/flights')
      .then(response => setFlights(response.data))
      .catch(error => console.error('Error fetching flight data:', error));
  }, []);

  return (
    <div>
      <h2>Flight Status</h2>
      <ul>
        {flights.map(flight => (
          <li key={flight.id}>
            Flight {flight.number}: {flight.status} (Gate: {flight.gate})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FlightStatus;
