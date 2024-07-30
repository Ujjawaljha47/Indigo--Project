// src/components/NotificationSettings.js
import React, { useState } from 'react';
import axios from 'axios';

const NotificationSettings = () => {
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post('/api/notifications', { email, phone })
      .then(response => alert('Settings saved!'))
      .catch(error => console.error('Error saving settings:', error));
  };

  return (
    <div>
      <h2>Notification Settings</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Email: </label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        </div>
        <div>
          <label>Phone: </label>
          <input type="tel" value={phone} onChange={(e) => setPhone(e.target.value)} />
        </div>
        <button type="submit">Save</button>
      </form>
    </div>
  );
};

export default NotificationSettings;
