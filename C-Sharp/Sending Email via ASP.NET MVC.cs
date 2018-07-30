 public ActionResult send_email(string name, string email,string Body,string Subject)
        {

            ViewBag.name = name;
            if (ModelState.IsValid)
            {
                var senderEmail = new MailAddress("senderEmailHere", "senderNameHere");
                var recEmail = new MailAddress(receiverEmailHere, "recieverNameHere");
                var password = "SenderPasswordHere";
                var body = Body;
                var sub = Subject;              
                var smpt = new SmtpClient
                {
                    Host = "smtp.gmail.com",
                    Port = 587,
                    EnableSsl = true,
                    DeliveryMethod = SmtpDeliveryMethod.Network,
                    UseDefaultCredentials = false,
                    Credentials = new NetworkCredential(senderEmail.Address, password)
                };
                using (var mess = new MailMessage(senderEmail, recEmail)
                {
                    Subject = sub,
                    Body = body
                }
                )
                {
                    smpt.Send(mess);
                }

            }
                return RedirectToAction("index");
        }
