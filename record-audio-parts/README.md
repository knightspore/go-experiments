# Outline

1. The program should chunk out 30s audio clips
  - [X] Get microphone input working
  - [ ] Split at 30s chunks
  - [ ] Convert to GoRoutine
2. Each of these should be passed to openai-whisper
  - [ ] Create module to use whisper-cli from Go
  - [ ] Create fifo queue for audio clips
  - [ ] Write transcriptions to cli and save to file
3. The output from the whisper detection should be printed in the terminal
