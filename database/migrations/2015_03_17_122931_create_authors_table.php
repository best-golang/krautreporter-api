<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateAuthorsTable extends Migration {

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('authors', function($table)
        {
            $table->integer('id')->unsigend();
            $table->string('name');
            $table->string('title');
            $table->string('image');
            $table->text('biography');
            $table->string('homepage')->nullable();
            $table->string('twitter')->nullable();
            $table->timestamps();
            
            $table->primary('id');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('authors');
    }

}