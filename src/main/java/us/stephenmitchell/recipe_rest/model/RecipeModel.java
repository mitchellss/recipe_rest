package us.stephenmitchell.recipe_rest.model;

import com.fasterxml.jackson.annotation.JsonFormat;

import java.util.Date;

import javax.persistence.*;

@Entity
@Table(name = "recipe")
public class RecipeModel {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name="ID")
    private Long id;

    @Column(name="TITLE")
    private String title;

    @Column(name="DATETIME")
    //@JsonFormat(pattern="yyyy-MM-dd HH:mm:ss")
    private String datetime;

    public RecipeModel() {
        super();
    }

    public RecipeModel(long id, String title, String datetime) {
        super();
        this.id = id;
        this.title = title;
        this.datetime = datetime;
    }

    public Long getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDatetime() {
        return datetime;
    }

    public void setDatetime(String datetime) {
        this.datetime = datetime;
    }
}
