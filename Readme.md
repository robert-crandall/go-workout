# Go Workout

This is a tool to describe workout cycles in Go.

Workouts can be exported to [Personal Trainer](https://apps.apple.com/us/app/personal-training-coach/id1325495597) format, 
and a human-readable format for comparisons.

## Structure

- `Lifts` are an exercise: Bench Press is a lift.
- `Sessions` are ways to describe the reps/sets/weights for a lift. 3x5 @ 85% for Bench Press is a session.
- `Programs` are ways to describe the complete set of an exercise goal. 3x5 for all lifts is the Starting Stength *program*.
