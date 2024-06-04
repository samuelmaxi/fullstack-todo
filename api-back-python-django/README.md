# API-BACK-PYTHON-DJANGO

## Path

| Path | Methods | Return (200) |
| --- | --- | --- |
| /api/v1/task/ | GET | Task list |
| /api/v1/task/search/{id}/ | GET | Search task |
| /api/v1/task/create/ | POST | Create new task |
| /api/v1/task/change/{int:id}/ | PUT | Change all attributes, but passing all attributes in the request body |
| /api/v1/task/edit/{int:id}/ | PATCH | Change the attributes separately, but without having to pass all the attributes in the request body |
| /api/v1/task/delete/{int:id}/ | DELETE | Delete the task |
