-- Table Name: Author

-- Column Name	Data Type	Description
-- author_id	Integer	Unique identifier for the author
-- name	Varchar(100)	Full name of the author
-- email	Varchar(100)	Email address of the author
-- affiliation	Varchar(100)	Affiliation or organization/college/university of the author
-- biography	Text	Biographical information about the author
-- website	Varchar(100)	Website or portfolio of the author

CREATE TABLE Author (
  author_id UUID UNIQUE NOT NULL,
  name VARCHAR(100),
  email VARCHAR(100),
  affiliation VARCHAR(100),
  biography TEXT,
  website VARCHAR(100)
);

--Sample Data that could be used to insert manaully if needed
INSERT INTO Author (author_id, name, email, affiliation, biography, website)
VALUES
  ('2b5e7f41-6b87-4a43-a23d-6b768babc123', 'John Smith', 'johnsmith@example.com', 'University of XYZ', 'John Smith is a professor of Computer Science...', 'https://www.johnsmith.com'),
  ('d3e4a16f-871d-4c5e-b523-6b82f1fcd987', 'Emily Johnson', 'emilyjohnson@example.com', 'Research Institute ABC', 'Emily Johnson is a researcher specializing in...', 'https://www.emilyjohnson.com'),
  ('5e49675f-6d3a-4e14-9b36-8b5b238c3e11', 'Michael Davis', 'michaeldavis@example.com', 'Company XYZ', 'Michael Davis is a software engineer with expertise in...', 'https://www.michaeldavis.com'),
  ('c1a83423-8e42-4fd3-aaf3-5a987e84033e', 'Sarah Thompson', 'sarahthompson@example.com', 'Independent Consultant', 'Sarah Thompson is an independent consultant with a focus on...', 'https://www.sarahthompson.com'),
  ('b0916fae-c26c-4dd4-9ebc-149e3d68e3a4', 'Robert Wilson', 'robertwilson@example.com', 'University of ABC', 'Robert Wilson is a professor of Economics...', 'https://www.robertwilson.com'),
  ('9a83b901-5209-4b53-980e-dcbbd56c7b59', 'Jessica Lee', 'jessicalee@example.com', 'Company ABC', 'Jessica Lee is a marketing specialist...', 'https://www.jessicalee.com'),
  ('ef05e974-78e0-40dd-9430-09059e40c0c5', 'David Brown', 'davidbrown@example.com', 'Research Center XYZ', 'David Brown is a leading scientist in the field of...', 'https://www.davidbrown.com'),
  ('c3e714a3-afcb-46a0-bbe6-3d1c695663db', 'Jennifer Clark', 'jenniferclark@example.com', 'University of XYZ', 'Jennifer Clark is an assistant professor in the Department of...', 'https://www.jenniferclark.com'),
  ('7246e07b-aae2-4e55-b091-1f0802a6d667', 'Andrew Evans', 'andrewevans@example.com', 'Company XYZ', 'Andrew Evans is a project manager...', 'https://www.andrewevans.com'),
  ('a55c4f3d-315b-4e5c-aae9-3e6016ef567d', 'Michelle Roberts', 'michelleroberts@example.com', 'Organization ABC', 'Michelle Roberts is a social worker with years of experience in...', 'https://www.michelleroberts.com');
