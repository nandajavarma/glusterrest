import restcurl


APP_ID = "MyApp"
APP_SECRET = "MySecret"
URL_ROOT = "https://hostname/v1"

http_status = {
    "200": "OK"
}


class Obj(object):
    def __init__(self, **kwargs):
        for k, v in kwargs.items():
            setattr(self, k, v)


class Api(object):
    def __init__(self, **kwargs):
        self.group = kwargs.get("group", "")
        self.title = kwargs.get("title", "")
        self.method = kwargs.get("method", "GET")
        self.url = kwargs.get("url", "")
        self.params = kwargs.get("params", [])
        self.resp = kwargs.get("resp", "")
        self.title_link = self.title.replace(" ", "_").lower()

    def print_params(self):
        params_iter = iter(self.params)
        op = ""
        try:
            val = params_iter.next()
            val.append("Required" if val[2] else "Optional")
            op += ("|Parameters:| {0} - `{1}` {3}(**{4}**)|\n".format(*val))
            for val in params_iter:
                val.append("Required" if val[2] else "Optional")
                op += ("|| {0} - `{1}` {3}(**{4}**)|\n".format(*val))
        except StopIteration:
            op += ("|Parameters:|None|\n")

        return op

    def print_index(self):
        return ("|{group}|{title}|{method}|[{url}]"
                "(#head_{title_link})|\n".format(
                    **self.__dict__))

    def print_detail(self):
        op = ""
        op += ("## <a name=\"head_{title_link}\">API {title}\n".format(
            **self.__dict__))

        op += ("|POST       |/peers|\n")
        op += ("|-----------|-----------------|\n")
        op += self.print_params()
        op += ("|Response: |HTTP/1.1 {0} {1}|\n".format(
            self.resp,
            http_status[str(self.resp)]))

        args = Obj(id=APP_ID, secret=APP_SECRET, method=self.method,
                   url=URL_ROOT + self.url,
                   show_headers=True)

        op += ("\nExample:\n\n")
        op += (restcurl.get_curl_cmd(args, space="    "))
        op += "\n"
        return op


class Apis(object):
    def __init__(self):
        self.apis = []

    def add(self, **kwargs):
        self.apis.append(Api(**kwargs))

    def doc(self):
        op = open("docs/0.md").read()
        op += "\n- APIs:\n\n"
        op += ("|Group          |Description           |Method |URL|\n")
        op += ("|---------------|----------------------|-------|---|\n")
        for api in self.apis:
            op += api.print_index()

        op += "\n"
        op += open("docs/1.md").read()
        op += "\n"
        for api in self.apis:
            op += api.print_detail()
            op += "\n"

        return op


api = Apis()
api.add(group="Peers", title="Peers Add", method="POST", url="/peers",
        params=[["hosts", "list", True, "List of Hosts"]],
        resp=200)

api.add(group="Peers", title="Peers Delete", method="DELETE", url="/peers",
        params=[["hosts", "list", True, "List of Hosts"]],
        resp=200)

with open("docs/API.md", "w") as f:
    f.write(api.doc())
