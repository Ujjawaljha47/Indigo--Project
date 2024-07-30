import React from 'react';
import './Dashboard.css';
import FlightStatus from './FlightStatus';
import NotificationSettings from './NotificationSettings';

const Dashboard = () => {
  return (
    <div className="dashboard-container">
      <h1>Flight Status and Notifications</h1>
      <div className="flight-status">
        <FlightStatus />
      </div>
      <div className="notification-settings">
        <NotificationSettings />
      </div>
    </div>
  );
};

export default Dashboard;
