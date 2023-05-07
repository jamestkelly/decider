import React from 'react';

function LandingPage() {
  return (
    <div className="authenticated-landing-container">
      <h1>Welcome, User!</h1>
      <p>You are authenticated and logged in. Here's some personalized content for you.</p>
      <section className="profile-section">
        <h2>Your Profile</h2>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin fringilla tempor augue vel bibendum.
          Quisque ullamcorper finibus mi, et vehicula elit venenatis at. Nam suscipit scelerisque eros a
          scelerisque. Sed facilisis, libero ac laoreet hendrerit, tortor libero consectetur urna, a dictum
          tellus ligula id augue.
        </p>
        <button className="profile-cta">Edit Profile</button>
      </section>
      <section className="dashboard-section">
        <h2>Your Dashboard</h2>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin fringilla tempor augue vel bibendum.
          Quisque ullamcorper finibus mi, et vehicula elit venenatis at. Nam suscipit scelerisque eros a
          scelerisque. Sed facilisis, libero ac laoreet hendrerit, tortor libero consectetur urna, a dictum
          tellus ligula id augue.
        </p>
        <button className="dashboard-cta">Go to Dashboard</button>
      </section>
    </div>
  );
}

export default LandingPage;
