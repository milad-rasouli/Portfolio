// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Milad75Rasouli/portfolio/frontend/views/layouts"
import "fmt"

func Contact(status string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = layouts.Base("Contact", fmt.Sprintf(`
<div class="b-content-divider b-content-divider-day b-content-vr col p-0">
    <div class="container h-100">
        <div class="h-100 align-hv-center">
            <div class="pt-3 pb-3 p-4 border border-light bg-opacity-75 text-start bg-dark col-9">
                    <h2>Get In Touch</h2>
                    <form class="w-100" action="/contact" method="post"> <!--TODO: CSRF protection-->
                        <label for="contact-subject">Subject:</label>
                        <input class="form-control" minlength="4" maxlength="100" type="text" name="subject" id="contact-subject" required>
                        <label for="contact-email">Email:</label>
                        <input class="form-control" type="email" name="email" id="contact-email" required>
                        <label for="contact-message">Message:</label>
                        <input class="form-control" type="text" value="its a csrf token" name="csrf_token" disabled hidden>
                        <textarea rows="5" style="resize: none;" class="form-control" type="text" name="message" id="contact-message" required></textarea>
                        <input class="btn btn-dark mt-2" type="submit" value="Submit"></input>
                    </form>
            </div>
        </div>
        </div>
    </div>
</div>
    <script>
    var notification = document.getElementById('notification');

    function ToggleNotification(){
        notification.style.display = 'block';
        setTimeout(function() {
        notification.style.display = 'none';
        }, 10000);
    }
    var status = %s;
    if (status="1"){
        console.log("status is "+ status)
        notification.textContent = "I got you message. I will reply it soon.";
        ToggleNotification();
    }else if(status =="2"){
        console.log("status is "+ status)
        notification.textContent = "Sadly, I could not get you message. you can contact with me by my email please navigate to About me to find it.";
        ToggleNotification();
    }else if(status =="3"){
        console.log("status is "+ status)
        notification.textContent = "Invalid fields! make sure your Subject and Email and Message meet the requirements.";
        ToggleNotification();
    }
    </script>
`, status)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}