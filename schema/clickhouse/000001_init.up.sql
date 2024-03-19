CREATE TABLE IF NOT EXISTS awesomedb.logs(
                                            Id INTEGER,
                                            ProjectId INTEGER,
                                            Name VARCHAR(255),
                                            Description VARCHAR(255),
                                            Priority INTEGER,
                                            Removed BOOLEAN,
                                            EventTime DATETIME
) ENGINE=MergeTree()
    ORDER BY (Id,ProjectId,Name);

INSERT INTO awesomedb.logs(Id, ProjectId, Name, Description, Priority, Removed, EventTime)
                            VALUES (0,0,'0','',0,false,now());
