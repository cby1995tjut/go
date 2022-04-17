<div class="timeline">
    {{ if .Tasks}} {{range .Tasks}}
        <div class="note">
            <p class="noteHeading">{{.Title}}</p>
            <hr>
            <p class="noteContent">{{.Content}}</p>
            </ul>
            </span>
        </div>
    {{end}} {{else}}
    <div class="note">
        <p class="noteHeading">No Tasks here</p>
        <p class="notefooter">
            Create new task<button class="floating-action-icon-add" > here </button> </p>
    </div>
{{end}}