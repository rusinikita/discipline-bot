# discipline-bot
Telegram bot that reminds you to track things and do things without overpush yourself

You don't need to do MORE. You need to do LESS, but REGULAR.

# Functions

### Tasks
- todo: /current_task Get current task in message
- todo: Marks task as done and send next task
- todo: Reminds about current task daily (task message, reminder task message with reply_to)
- todo: Reminds about conditional task at some time (buy something after work, do at home)

### Tracking
- todo: /track_<tracker_name> <value> Writes value for tracker
- todo: Reminds about important things to track daily
- todo: /track_stats Send some stats about important trackers

### Routines
- todo: /routine_<routine_name> Start passing come checklist (track something, do something)
  - design: message with action buttons
    - true/false done button
    - <track_short_name> <value> in user message will be interpreted
- todo: Reminds about important routine daily