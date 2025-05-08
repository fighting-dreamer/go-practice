## Todo App

Two main entities :

1. List :
   id
   name
   created-by
   created-on
   task-count

2. Task
   id
   title
   list#ref
   status
   labels
   created-by
   created-on
   due-by
   completed-on

3. ListAggregations :
   list#ref
   task-count
   task-count-by-label
   task-count-by-status

Two components :

1. backend
2. cli/user-facing

---

1. user have a api-key to configure in cli to use.
2. user can crud on list
3. user can crud on task
4. user can update for a task :
   - you can have dynamically list of attributes
